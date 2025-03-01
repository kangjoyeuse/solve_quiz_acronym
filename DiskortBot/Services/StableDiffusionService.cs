namespace DiskortBot.Services;

using System.Text;
using System.Text.Json;
using System.Net.Http.Json;
using Microsoft.Extensions.Logging;

public class StableDiffusionService : IDisposable
{
    private readonly HttpClient _httpClient;
    private readonly ILogger<StableDiffusionService> _logger;

    public StableDiffusionService(
        HttpClient httpClient,
        ILogger<StableDiffusionService> logger)
    {
        _logger = logger;
        _httpClient = httpClient;
    }

    public record StableDiffusionConfig(
        string Prompt,
        string NegativePrompt,
        int Steps = 25,
        double CfgScale = 7.0,
        int Width = 1024,
        int Height = 1024,
        int Seed = -1,
        string SamplerName = "Euler a");

    private record GenerationResponse(string[] Images, string Info);

    public record StableDiffusionResultInfo(int Seed);

    public record StableDiffusionResult(byte[] Image, StableDiffusionResultInfo Info);

    public async Task<StableDiffusionResult> GenerateAsync(StableDiffusionConfig config)
    {
        var opt = new JsonSerializerOptions
        {
            PropertyNamingPolicy = JsonNamingPolicy.SnakeCaseLower
        };
        var payload = JsonSerializer.Serialize(config, opt);

        var content = new StringContent(payload, Encoding.UTF8, "application/json");
        content.Headers.ContentLength = payload.Length;

        var response = await _httpClient.PostAsync("/sdapi/v1/txt2img", content);
        response.EnsureSuccessStatusCode();

        var result = await response.Content.ReadFromJsonAsync<GenerationResponse>();

        return new(
            Convert.FromBase64String(result!.Images[0]),
            JsonSerializer.Deserialize<StableDiffusionResultInfo>(result!.Info, opt)!
        );
    }

    public void Dispose()
    {
        _httpClient?.Dispose();
    }
}
