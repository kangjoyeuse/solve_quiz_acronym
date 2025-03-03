import pyautogui
import time
import random

# Spam message
message = "Hello, I am a bot. I am spamming this message."

# Delay between each message
delay = 0.5

# Number of messages to spam
num_messages = 100

# Start spamming
for i in range(num_messages):
    pyautogui.typewrite(message)
    pyautogui.press('enter')
    time.sleep(delay)
    print(f"Message {i + 1} sent.")
    time.sleep(random.uniform(0.5, 1.5))  # Random delay between 0.5 and 1.5 seconds

print("Spamming complete.")
