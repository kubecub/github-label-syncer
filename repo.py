# Copyright © 2023 KubeCub open source community. All rights reserved.
# Licensed under the MIT License (the "License");
# you may not use this file except in compliance with the License.

import requests

url = "https://api.github.com/orgs/kubecub/repos"

response = requests.get(url)
data = response.json()

for repo in data:
    html_url = repo["html_url"]
    repo_name = html_url.replace("https://github.com/", "")
    print(repo_name)
