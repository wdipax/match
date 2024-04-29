# Speed Dating Telegram Bot

## The Flow

The bot awaits for a message from admin. 
Then it responds with two links. One for boys and one for girls.
Admin should send those links in private messages to every guest.
Guests can join groups using these links.

Bot will display control buttons for the admin. 
Admin can guide the event by using those buttons.

After the guests know each other the admin starts voting.
Guests vote for each other.
Admin can see the statistics on the number of guests who voted.
Admin has to end voting when ready.
After this, all guests will receive contacts of their matches in Telegram.

## Build

```sh
docker build -t match:latest --build-arg LANG=en .
```

## Run

```sh
docker run --rm -e TELEGRAM_BOT_TOKEN=$YOUR_TG_BOT_TOKEN -e ADMIN_USER_NAME=$YOUR_TG_USERNAME -e DEBUG=false match:latest
```
