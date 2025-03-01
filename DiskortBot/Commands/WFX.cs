namespace DiskortBot.Commands;

using System.Text.RegularExpressions;
using Discord;
using Discord.Interactions;

using DiskortBot.Services;

[Group("wfx", "WFX commands")]
public class WFX : InteractionModuleBase<SocketInteractionContext>
{
    private readonly StableDiffusionService _sdService;

    public WFX(StableDiffusionService sdService)
    {
        _sdService = sdService;
    }

    [SlashCommand("generate", "Generate image")]
    public async Task Generate(
        [Summary("prompt", "Prompt Input")] string prompt,
        [Summary("negative", "Negative Prompt Input")] string negativePrompt = "lowres, bad anatomy, bad hands, text, error, missing finger, extra digits, fewer digits, cropped, worst quality, low quality, low score, bad score, average score, signature, watermark, username, blurry",
        [Summary("ratio", "Ratio Image (ex: 1024x1024)")] string ratio = "1024x1024",
        [Summary("steps", "Steps (min: 1, max: 50)")] int steps = 25,
        [Summary("cfg", "CFG Scale")] double cfg = 7.0)
    {
        await DeferAsync();

        try
        {
            var dimension = GetDimension(ratio);
            var image = await _sdService.GenerateAsync(new(
                prompt,
                negativePrompt,
                Width: dimension.Width,
                Height: dimension.Height,
                Steps: steps,
                CfgScale: cfg
            ));

            await FollowupWithFileAsync(
                new FileAttachment(new MemoryStream(image.Image), "generated.png"),
                text: $"🎨 {steps} steps, {cfg} cfg, {dimension.Width}x{dimension.Height} ratio, {image.Info.Seed} seed"
            );
        }
        catch (Exception error)
        {
            await FollowupAsync($"Exception: {error.Message}");
        }
    }

    private record Dimension(int Width, int Height);

    private Dimension GetDimension(string ratio)
    {
        Match patern = Regex.Match(ratio, @"(\d+)x(\d+)");

        if (!patern.Success)
        {
            throw new Exception("Width/Height must be a Int!");
        }

        return new(int.Parse(patern.Groups[1].Value), int.Parse(patern.Groups[2].Value));
    }
}
