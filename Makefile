push:
	git pull
	git add .
	git commit -m "update" || true
	git push origin hbj-all || true
	