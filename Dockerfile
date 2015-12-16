FROM scratch
ADD pathfinder /
ADD docs /docs
EXPOSE 8080
CMD ["/pathfinder"]

