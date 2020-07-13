#!/bin/sh
source .env
curl https://api.telegram.org/bot$TELEGRAM_BOT_TOKEN/setWebhook?url=$TELEGRAM_POST_ENDPOINT
