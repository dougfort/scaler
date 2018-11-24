FROM alpine

COPY ./settings.toml /settings.toml
COPY ./scaler /scaler
COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 10000
EXPOSE 8080