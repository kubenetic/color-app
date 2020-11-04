# color-app

Az alkalmazás Kubernetes objektumok létrehozása és használata (Pod, ConfigMap) skálázódási és rolling-update valamint history kezelésének bemutatása végett született. 

## Alkalmazás konfigurálása

Az alkalmazásnak az alábbi két környezeti változó adható át amik a megjelenő szöveg színét és az oldal háttér színét állítják be:

* APP_FONT_COLOR (alapértelmezetten '_white_')
* APP_BG_COLOR (alapértelmezetten '_blue_')

## Alkalmazás image létrehozása
```bash
buildah bud --no-cache \
    --rm \
    --tag kubenetic/color-app:latest 
```

## Alkalmazás futtatása

```bash
podman run --detach \
    --env APP_FONT_COLOR='white' \
    --env APP_BG_COLOR='blue' \
    --publish 8080:8080 \
    kubenetic/color-app:latest
```