FROM openjdk:17.0.1-jdk-slim

ENV LANG C.UTF-8
ENV TZ Asiz/Tokyo

ENV GCLOUD_VERSION 371.0.0

RUN apt update
RUN apt install -y --no-install-recommends \
  bash \
  curl \
  python \
  tar

RUN curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-${GCLOUD_VERSION}-linux-x86_64.tar.gz
RUN tar -xvzf "google-cloud-sdk-${GCLOUD_VERSION}-linux-x86_64.tar.gz" \
  && rm -f "google-cloud-sdk-${GCLOUD_VERSION}-linux-x86_64.tar.gz" \
  && ./google-cloud-sdk/install.sh --usage-reporting true --rc-path ~/.bashrc --command-completion true --path-update true \
  && ./google-cloud-sdk/bin/gcloud components install beta pubsub-emulator --quiet

EXPOSE 8085
CMD ["./google-cloud-sdk/bin/gcloud", "beta", "emulators", "pubsub", "start", "--host-port=0.0.0.0:8085"]
