echo "$(date) Repo Quickstart" && \
    git init && \
    git config user.email "$GH_USER_EMAIL" && \
    git config user.name "$GH_AUTHOR_NAME" && \
    git add . && \
    git commit -m "init" && \
    git branch -M main && \
    git remote add origin $GH_REPO_URL && \
    git push -u origin main