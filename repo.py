import requests

url = "https://api.github.com/orgs/kubecub/repos"

response = requests.get(url)
data = response.json()

for repo in data:
    html_url = repo["html_url"]
    repo_name = html_url.replace("https://github.com/", "")
    print(repo_name)
