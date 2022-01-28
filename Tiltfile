allow_k8s_contexts('colima')

docker_build(
    'docker-hosted.artifactory.tcsbank.ru/devplatform/kratos',
    '.',
    dockerfile='./build/Dockerfile-build',
    build_args={'GITLAB_TOKEN': 'VkHzEBp_BAPEx8LSFjty'}
)

k8s_yaml(kustomize('/Users/a.koval/go/src/gitlab.tcsbank.ru/devplatform/local-development/iam/'))

k8s_resource('kratos') #, port_forwards=3001)
