# Developer take home technical challenge

In terms of time, the challenge is intended to be completed over a few hours - it is not meant to consume a weekend. 
We recognize that personal and professional commitments exist. If there are any limitations in the completed submission due to time constraints, please mention these in the README.

The intent of the challenge is to gain insights on problem-solving and coding.
Documentation should be provided for compiling and running the solution.
Evaluation will be based on accuracy, readability, design choices, and 
completeness.

The solution should be made available as a public GitHub
repository.

## Problem description

Pretend we have a tool that reports information from a machine to our service in
JSON format. We wish to capture this information in a server for easy viewing.

For this project we want to create a RESTful API that accepts the JSON 
payloads <sup>1</sup> and stores it in a datastore<sup>2</sup>. The RESTful API
should  also have a GET request that returns all entries in the store.

The endpoints, project structure, API documentation and library options are up
to you. The project should be unit tested to prove that your project works.


## Examples

### Basic Post Request
Here's a basic curl request that you would expect to succeed:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "machineId": 12345,
    "stats": {
        "cpuTemp": 90,
        "fanSpeed": 400,
        "HDDSpace": 800
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
}' \
  https://urldefense.proofpoint.com/v2/url?u=http-3A__localhost-3A4000_&d=DwIGaQ&c=BioHiDP8cpFFEOWoiyRJQw&r=8lDpqcXRvcOUrXln89I8_aXex2o2ePZlwX18XSAF-4M&m=_zzicL0bPOGP54q1qYaKz5bEC-7QdNfJX9aNB1sywHpILpf6o8jAzV2980Oj_hh0&s=zWjEMl8HkC25oHxe5x0og43Czrt_xFOMnuKLtzb1hiQ&e= <path>/<to>/<your>/<endpoint>
```
We would expect to see a `201` response code. The body can be what you determine
to be appropriate.

### Basic GET request
On Fetching data, here's an example assuming that the previous POST request
returned is the only record in the data store:

```
curl https://urldefense.proofpoint.com/v2/url?u=http-3A__localhost-3A4000_&d=DwIGaQ&c=BioHiDP8cpFFEOWoiyRJQw&r=8lDpqcXRvcOUrXln89I8_aXex2o2ePZlwX18XSAF-4M&m=_zzicL0bPOGP54q1qYaKz5bEC-7QdNfJX9aNB1sywHpILpf6o8jAzV2980Oj_hh0&s=zWjEMl8HkC25oHxe5x0og43Czrt_xFOMnuKLtzb1hiQ&e= <path>/<to>/<your>/<endpoint>
```
And the response may look something like:
```
[
    {
        "id": "5a0ff41d-fe3d-4e54-ad0e-e67a8401adb3",
        "machineId": 12345,
        "stats": {
            "cpuTemp": 90,
            "fanSpeed": 400,
            "HDDSpace": 800
        },
        "lastLoggedIn": "admin/Paul",
        "sysTime": "2022-04-23T18:25:43.511Z"
    }
]
```

<sup>1</sup> json payload example:

```
{
    "machineId": 12345,
    "stats": {
        "cpuTemp": 90,
        "fanSpeed": 400,
        "HDDSpace": 800
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
},
{
    "machineId": 4444,
    "stats": {
        "cpuTemp": 78,
        "fanSpeed": 500,
        "HDDSpace": 100
        "internalTemp": 23
    },
    "lastLoggedIn": "admin/Ian",
    "sysTime": "2022-04-21T19:25:43.219Z"
},
{
    "machineId": 61616,
    "stats": {
        "cpuTemp": "78c",
        "fanSpeed": 500,
        "HDDSpace": 100
        "internalTemp": 23
    },
    "lastLoggedIn": "admin/Tim",
    "sysTime": "Wed 2021-07-28 14:16:27"
}
```

<sup>2</sup> Feel free to use an in-memory data store such as a struct or map if
you want to save time. We don't expect a database to be needed unless you're
more comfortable writing the project with one.