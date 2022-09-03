FROM alpine:latest

#8 run command on a new docker image
RUN mkdir /app

#9 build from aiuth app and copy to /app
#COPY --from=builder /app/authrApp /app
COPY authApp /app

#10 execute command
CMD [ "/app/authApp"]