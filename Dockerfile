FROM centos
WORKDIR /app
COPY . /app
EXPOSE 80
VOLUME /app/blog/log
ENV GIN_MODE=release
CMD sleep 10 && ./blogger
# 等待10秒让数据库启动