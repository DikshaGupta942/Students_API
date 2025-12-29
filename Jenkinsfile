pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'gdiksha942/students-api'
    }

    stages {
        stage('Checkout Code') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/DikshaGupta942/Students_API.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE:latest .'
            }
        }

        stage('Docker Hub Login') {
            steps {
                withCredentials([usernamePassword(
                    credentialsId: 'dockerhub-creds',
                    usernameVariable: 'DOCKER_USER',
                    passwordVariable: 'DOCKER_PASS'
                )]) {
                    sh 'echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin'
                }
            }
        }

        stage('Push Image') {
            steps {
                sh 'docker push $DOCKER_IMAGE:latest'
            }
        }

        stage('Deploy Container') {
            steps {
                sh '''
                echo "Stopping old container if exists..."
                docker ps -q --filter "publish=8081" | xargs -r docker stop
                docker ps -aq --filter "publish=8081" | xargs -r docker rm

                echo "Pulling latest image..."
                docker pull gdiksha942/students-api:latest

                echo "Starting new container..."
                docker run -d \
                    -p 8081:8082 \
                    --name students-api \
                    gdiksha942/students-api:latest
                '''
            }
        }
    }

    post {
        success {
            echo 'Docker image built, pushed and deployed successfully!'
        }
        failure {
            echo 'Pipeline failed.'
        }
    }
}
