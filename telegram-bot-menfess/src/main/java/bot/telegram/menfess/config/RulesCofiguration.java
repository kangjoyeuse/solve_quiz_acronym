package bot.telegram.menfess.config;

import lombok.Data;

@Data
public class RulesCofiguration {


    int limitFreeUsers = 3;
    int payAmountAfterLimit = 100;
    boolean forceShowUsername = true;
    long getChannelId = -1002161694809L;

}
