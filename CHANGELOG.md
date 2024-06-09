# CHANGELOG

## 0.0.1
Initial release with the capability to continuously log messages to `/var/log/random.log` until the program is terminated.
- Includes 3 different possible log messages
- Supports 3 log levels: **INFO**, **WARNING**, and **ERROR**

Example log entries:

```plaintext
2024/06/05 21:43:14 [WARNING] Another random log message
2024/06/05 21:43:17 [INFO] This is a random log message
2024/06/05 21:43:19 [INFO] Random log message number three
2024/06/05 21:45:01 [ERROR] Random log message number three
2024/06/05 21:45:03 [INFO] Random log message number three
2024/06/05 21:45:04 [WARNING] Another random log message
2024/06/05 21:45:06 [ERROR] Random log message number three
```

# 0.2
1. Enhanced log message generation with a `randomSentence` function, creating random messages composed of **subjects** + **verbs** + **objects** + **adverbs**
    - Each part contains 50 unique items, resulting in 6,250,000 possible combinations.

Example log entries:

```plaintext
2024/06/05 21:50:32 [INFO] The spider ventures into a burrow boldly.
2024/06/05 21:50:35 [INFO] A whale swirls in a boulder hurriedly.
2024/06/05 21:50:36 [ERROR] A dolphin discovers a cliff loudly.
2024/06/05 21:50:37 [ERROR] A seagull runs to a tree recklessly.
2024/06/05 21:50:38 [ERROR] The seal dives down a cave earnestly.
2024/06/05 21:50:41 [INFO] The seal catches the tide patiently.
2024/06/05 21:50:44 [WARNING] The bear barks at the stars nonchalantly.
2024/06/05 21:50:47 [WARNING] The rabbit catches the rock suddenly.
2024/06/05 21:50:49 [ERROR] A squirrel sneaks up on the sand briskly.
2024/06/05 21:50:51 [WARNING] The panda swirls in the lagoon relentlessly.
```

2. Put in place `CHANGELOG.md` and tagging
3. Reworked `readme.md` accordingly