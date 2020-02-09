# fritzbox clourflare dyndns

Simply proxies fritzbox dydns update requests to cloudflare

## Usage:

### Pre-requests
* kubernetes cluster
* Cloudflare ApiKey
* Cloudflare EMail 
* Cloudflare zone is created
* Cloudflare dns entry is created, because this will only be updated

### Installing with helm chart
First of all you need to install this to some kubernetes cluster

copy this to a file called values.yaml
```yaml
ingress:
  enabled: true
  annotations:
    kubernetes.io/ingress.class: nginx
  hosts:
    - host: chart-example.local
      paths: ["/"]
```

```bash
git clone https://github.com/JSchlarb/fritzbox-cloudflare
cd fritzbox-cloudflare/charts
helm install fritxbox-cloudflare ./fritxbox-cloudflare -f values.yaml
```

### Setup fritzbox

![fritxbox ui][docs/fritzbox-ui.png]

Setup your fritzbox 
* DynDns-Anbieter: `Benutzerdefiniert`
* UpdateUrl: `http://chart-example.local/api/dyndns?hostname=<domain>&ipAddress=<ipaddr>&zoneId=$$ZONEID$$&dnsId=$$DNSID$$` (* set $$ZONEID$$ and $$DNSID$$ with the correct values)
* Benutzername: `cloudflare email`
* Kennwort: `cloudflare apikey`

## TBD:
* improve docs
