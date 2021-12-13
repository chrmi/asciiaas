FROM scratch
EXPOSE 8099
ENV APIPORT=8099
ADD ./build/asciiaas ./
CMD ["./asciiaas"]
