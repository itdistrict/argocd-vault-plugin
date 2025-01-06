FROM alpine
RUN apk add --no-cache kubectl kustomize helm
COPY argocd-vault-plugin /usr/local/bin/argocd-vault-plugin
RUN chmod +x /usr/local/bin/argocd-vault-plugin