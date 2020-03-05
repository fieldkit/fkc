@Library('conservify') _

timestamps {
    node () {
        stage ('git') {
			git branch: 'master', url: 'https://github.com/fieldkit/fkc.git'
        }

        stage ('clean') {
            sh "make clean"
        }

        stage ('build') {
            withEnv(["PATH+GOLANG=${tool 'golang-amd64'}/bin"]) {
				withEnv(["PATH+GOHOME=${HOME}/go/bin"]) {
					sh "GOOS=darwin GOARCH=amd64 BINARY=fkc-darwin-amd64 make"
					sh "GOOS=linux GOARCH=amd64 BINARY=fkc-linux-amd64 make"
                }
            }
		}

        stage ('archive') {
			archiveArtifacts artifacts: "build/fkc*"
		}
    }
}
