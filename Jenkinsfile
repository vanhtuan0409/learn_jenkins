def app

pipeline {
    agent any

    stages {
        stage('build') {
            steps {
                echo "Starting to build docker image"
                script {
                    app = docker.build("sample_go_app", "./sample_project")
                }
            }
        }

        stage('test') {
            steps {
                script {
                    app.inside {
                        sh 'app > /var/log/app.txt &'
                        sh 'curl -s localhost:7777'
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
