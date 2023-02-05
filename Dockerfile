FROM alpine
ADD user_linux /user
ENTRYPOINT [ "/user" ]
