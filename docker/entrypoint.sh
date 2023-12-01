#!/usr/bin/env sh
/usr/app/dlv --listen=:${DELVE_DEBUG_PORT} --headless=true --api-version=2 --accept-multiclient exec /usr/app/go-code-examples
