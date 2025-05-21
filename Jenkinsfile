pipeline {
    agent any

    tools {
        go "1.24.3"
    }

    environment {
        GO111MODULE = "on"
        APP_NAME = "hello-world-api"
        APP_VERSION = "1.0.0"
        DOCKER_IMAGE = "hello-world-api:${APP_VERSION}"
        DOCKER_REGISTRY = "docker.io"
        GIT_URL = "https://github.com/huugiii/${APP_NAME}.git"
        GIT_BRANCH = "main"
    }

    stages {
        stage('Build Go Application') {
            steps {
                echo "========Building Go Application========"
                git url: GIT_URL, branch: GIT_BRANCH
                sh "go build -o hello-world-api cmd/app/main.go"
            }
            post {  
                success {
                    echo "========Go build successful========"
                    sh "ls -l"
                }
                failure {
                    echo "========Go build failed========"
                }
            }
        }

        stage('Build Docker Image') {
            steps {
                echo "========Building Docker Image========"
                sh "docker build -t ${DOCKER_IMAGE} . --build-arg APP_NAME=${APP_NAME}"
            }
            post {
                success {
                    echo "========Docker build successful========"
                }
                failure {
                    echo "========Docker build failed========"
                }
            }
        }

        stage('Push Docker Image') {
            steps {
                echo "========Pushing Docker Image========"
                withCredentials([usernamePassword(credentialsId: 'docker-credentials', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASSWORD')]) {
                    sh """
                        docker login ${DOCKER_REGISTRY} -u ${DOCKER_USER} -p ${DOCKER_PASSWORD}
                        docker tag ${DOCKER_IMAGE} ${DOCKER_REGISTRY}/${DOCKER_IMAGE}
                        docker push ${DOCKER_REGISTRY}/${DOCKER_IMAGE}
                    """
                }
            }
            post {
                success {
                    echo "========Docker push successful========"
                }
                failure {
                    echo "========Docker push failed========"
                }
            }
        }

        stage('Deploy') {
            steps {
                echo "========Deploying Application========"
                sh """
                    docker pull ${DOCKER_REGISTRY}/${DOCKER_IMAGE}
                    docker stop ${APP_NAME} || true
                    docker rm ${APP_NAME} || true
                    docker run -d --name ${APP_NAME} -p 8080:8080 ${DOCKER_REGISTRY}/${DOCKER_IMAGE}
                """
            }
            post {
                success {
                    echo "========Deployment successful========"
                }
                failure {
                    echo "========Deployment failed========"
                }
            }
        }
    }

    post {
        always {
            echo "========Pipeline completed========"
        }
    }
}