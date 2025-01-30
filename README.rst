.. raw:: html

.. image:: docs/source/\_static/youtube.png
:width: 100%
:alt: youtube link
:align: center
:target: https://www.youtube.com/embed/OFCKEXGSE3A

Device Connect (Fork)
=====================

This is a fork of `Google Cloud Platform's Device Connect https://github.com/GoogleCloudPlatform/deviceconnect`.  
It is designed to **consume Fitbit API endpoints** and **store results in BigQuery** as a backend-only service.

🚀 Features
-----------
- **Forked from Google Cloud Platform's `deviceconnect https://github.com/GoogleCloudPlatform/deviceconnect`**
- **Backend-only app** (OIDC flow removed, only exposing API endpoints)
- **Extended Fitbit API support** with additional endpoints
- **Enhanced documentation**
- **Docker support** for easy deployment
- **Deployed on Google Cloud Run**


🛠 Setup & Usage
----------------

1️⃣ **Environment Variables**
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
Make sure to configure your `.env` file before running the service.  
This file should contain all necessary environment variables required for authentication and configuration. Check the `env.example` file for reference.

2️⃣ **Building the Docker Image**
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
If you're building from a **Mac** and deploying to **Cloud Run**, ensure the correct platform is specified (`linux/amd64`).

.. code-block:: bash

    docker buildx build --platform linux/amd64 -t <GCP_REGION>-docker.pkg.dev/<GCP_PROJECT>/<ARTIFACT_REGISTRY>/<IMAGE_NAME>:<TAG> .

- ``<GCP_REGION>`` → Your Google Cloud region (e.g., ``us-central1``)
- ``<GCP_PROJECT>`` → Your Google Cloud project ID
- ``<ARTIFACT_REGISTRY>`` → Your Artifact Registry repository
- ``<IMAGE_NAME>`` → Your desired Docker image name
- ``<TAG>`` → Version tag (e.g., ``v1``, ``latest``)

---

3️⃣ **Running the Container Locally**
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
To run the container locally with a GCP service account, use the following command:

.. code-block:: bash

    docker run --platform linux/amd64 \
        --env-file .env \
        -v <PATH_TO_SERVICE_ACCOUNT_JSON>:/app/credentials.json \
        -p <PORT>:<PORT> \
        <IMAGE_ID>

- ``<PATH_TO_SERVICE_ACCOUNT_JSON>`` → Full path to your Google Cloud service account JSON key
- ``<PORT>`` → The port number your app runs on (e.g., ``3000``)
- ``<IMAGE_ID>`` → Docker image ID or name

The service accound should have the following roles:

- Secret Manager Secret Accessor
- Artifact Registry Reader
- Service Account Token Creator
- Firebase Permissions
- BigQuery Permissions
