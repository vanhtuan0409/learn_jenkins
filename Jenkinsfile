def app

pipeline {
    agent any

    environment {
        IMAGE_NAME = "vanhtuan/sample_go_app"
    }

    stages {
        stage('build') {
            steps {
                echo "Starting to build docker image"
                script {
                    app = docker.build("${env.IMAGE_NAME}", "./sample_project")
                }
            }
        }

        stage('test') {
            steps {
                script {
                    app.inside {
                        sh 'app &'
                        sh 'curl -s localhost:7777'
                    }
                }
            }
        }

        stage('input') {
            steps {
                script {
                    milestone 1
                    timeout(time: 2, unit: 'MINUTES') {
                        input "Should publish to docker hub?"
                    }
                    milestone 2
                }
            }
        }

        stage('publish') {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
                
            }
        }
    }

    post {
        success {
            echo "Success"
        }
        failure {
            echo "Failure"
        }
    }
}
