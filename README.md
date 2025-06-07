## URL SHORTENER

### Requirements
- Functional
    1. user should be able to get shortest url back on passing long_url
    2. when user passed short url it will reditrected to long url(original url)
    3. High available, fault tolerant and scalable
    4. Service should collect metrics like most clicked links
    5. Once a shortened link is generated it should stay in system for lifetime

- API DESIGN
    - POST /v1/shorten-url
        - URL (userName string , longURL string)
    - GET API /<