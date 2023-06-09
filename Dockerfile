FROM alpine
COPY argocd-vault-plugin /usr/local/bin/argocd-vault-plugin
RUN chmod +x /usr/local/bin/argocd-vault-plugin