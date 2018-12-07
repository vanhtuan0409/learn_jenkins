pipeline {
    agent { docker { image 'golang:alpine' } }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}
