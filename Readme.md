# Simple App per esempi dei corsi Red Hat

Questa app è una semplice e banale app scritta in Go da usare come esempio nei corsi DO{180,288,280,285,295} di Red Hat.


## Come usarla

Per usare l'applicazione basta deployarla su OpenShift con il comando:

`oc new-app https://github.com/elroncio/simple-app-rht`

Può essere deployata anche andando a customizzare il processo di build s2i con il branc `s2i`:

`oc new-app https://github.com/elroncio/simple-app#s2i`

## Endpoint esposti

L'applicazione espone i seguenti endpoint:

- `/static`
- `/healtz`
- `/cat`
- `/cats/[number]`
- `/hello`
- `/msg`
- `/help`


Le risposte ritornate sono in json, quindi è meglio parsarle con `jq` o formatter simili.