pipeline {
    agent {
        label 'docker-agent-alpine'
    }

    parameters {
        string(name: 'APP_NAME', defaultValue: "hello-world-api", description: "The name of the application")
        string(name: 'BUILD_VERSION', defaultValue: "1.0.0", description: "The version of the application")
        string(name: 'SOURCE_BRANCH', defaultValue: "main", description: "The branch of the git repository")
    }

    environment {
        GO111MODULE = "on"
        GIT_URL = "https://github.com/huugiii/${APP_NAME}.git"
        DOCKER_IMAGE = "huugiii/${APP_NAME}:${BUILD_VERSION}"
        DOCKER_REGISTRY = "docker.io"
    }

    triggers {
        pollSCM('H H(0-2) * * 1')
    }

    stages {
        stage('Validate') {
            steps {
                script {
                    if (APP_NAME == null || APP_NAME == "") {
                        error("Invalid app name")
                    }
                    if (BUILD_VERSION == null || BUILD_VERSION == "") {
                        error("Invalid build version")
                    }
                    if (SOURCE_BRANCH == null || SOURCE_BRANCH == "") {
                        error("Invalid source branch")
                    }
                }
            }
        }

        stage('Checkout') {
            steps {
                echo "Checking out ${GIT_URL} [${SOURCE_BRANCH}]"
                checkout scmGit(branches: [[name: "*/${SOURCE_BRANCH}"]], extensions: [], userRemoteConfigs: [[url: GIT_URL]])
            }
        }

        stage('Build & Push Image') {
            steps {
                script {
                    echo "Building image ${DOCKER_REGISTRY}/${DOCKER_IMAGE}"
                    dockerImage = docker.build("${DOCKER_REGISTRY}/${DOCKER_IMAGE}")

                    echo "Pushing image to ${DOCKER_REGISTRY} [${major}.${minor}.${patch}, ${major}.${minor}, ${major}, latest]"
                    major = BUILD_VERSION.split(".")[0]
                    minor = BUILD_VERSION.split(".")[1]
                    patch = BUILD_VERSION.split(".")[2]
                    docker.withRegistry("${DOCKER_REGISTRY}", "docker-credentials") {
                        dockerImage.push("${major}.${minor}.${patch}")
                        dockerImage.push("${major}.${minor}")
                        dockerImage.push("${major}")
                        dockerImage.push("latest")
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                sh """
                    docker pull ${DOCKER_REGISTRY}/${DOCKER_IMAGE}
                    docker stop ${APP_NAME} || true
                    docker rm ${APP_NAME} || true
                    docker run -d --name ${APP_NAME} -p 8080:8080 ${DOCKER_IMAGE}
                """
            }
        }
    }

    post {
        always {
            echo "========Pipeline completed========"
        }
    }
}