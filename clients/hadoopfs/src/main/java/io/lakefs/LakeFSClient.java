package io.lakefs;

import io.lakefs.auth.LakeFSTokenProvider;
import io.lakefs.auth.LakeFSTokenProviderFactory;
import io.lakefs.clients.sdk.*;
import io.lakefs.clients.sdk.auth.HttpBasicAuth;
import io.lakefs.clients.sdk.auth.HttpBearerAuth;
import org.apache.hadoop.conf.Configuration;

import java.io.IOException;

import static io.lakefs.auth.LakeFSTokenProviderFactory.newLakeFSTokenProvider;

/**
 * Provides access to lakeFS API using client library.
 * This class uses the configuration to initialize API client and instance per API interface we expose.
 */
public class LakeFSClient {
    private static final String BASIC_AUTH = "basic_auth";
    private static final String JWT_TOKEN_AUTH = "jwt_token";

    private final ObjectsApi objectsApi;
    private final StagingApi stagingApi;
    private final RepositoriesApi repositoriesApi;
    private final BranchesApi branchesApi;
    private final ConfigApi configApi;
    private final InternalApi internalApi;

    public LakeFSClient(String scheme, Configuration conf) throws IOException {
        String authProvider = FSConfiguration.get(conf, scheme, Constants.LAKEFS_AUTH_PROVIDER_KEY_SUFFIX, LakeFSClient.BASIC_AUTH);
        ApiClient apiClient;

        if (authProvider == BASIC_AUTH) {
            String accessKey = FSConfiguration.get(conf, scheme, Constants.ACCESS_KEY_KEY_SUFFIX);
            if (accessKey == null) {
                throw new IOException("Missing lakeFS access key");
            }

            String secretKey = FSConfiguration.get(conf, scheme, Constants.SECRET_KEY_KEY_SUFFIX);
            if (secretKey == null) {
                throw new IOException("Missing lakeFS secret key");
            }
            apiClient = newApiClientNoAuth(scheme, conf);
            HttpBasicAuth basicAuth = (HttpBasicAuth) apiClient.getAuthentication(BASIC_AUTH);
            basicAuth.setUsername(accessKey);
            basicAuth.setPassword(secretKey);
        } else {
            // TODO(isan) depends on missing functionality PR https://github.com/treeverse/lakeFS/pull/7578 being merged.
            //  once merged, we can use the following code to get the token
            throw new IOException("Unsupported auth provider: " + authProvider + ". Only basic_auth is supported at the moment.");
//            LakeFSTokenProvider tokenProvider = newLakeFSTokenProvider(scheme, conf);
//            String jwt = tokenProvider.getToken();
//            apiClient = newApiClientNoAuth(scheme, conf);
//            HttpBearerAuth tokenAuth = (HttpBearerAuth)apiClient.getAuthentication(JWT_TOKEN_AUTH);
//            tokenAuth.setBearerToken(jwt);
        }

        this.objectsApi = new ObjectsApi(apiClient);
        this.stagingApi = new StagingApi(apiClient);
        this.repositoriesApi = new RepositoriesApi(apiClient);
        this.branchesApi = new BranchesApi(apiClient);
        this.configApi = new ConfigApi(apiClient);
        this.internalApi = new InternalApi(apiClient);
    }

    ApiClient newApiClientNoAuth(String scheme, Configuration conf) {

        ApiClient apiClient = io.lakefs.clients.sdk.Configuration.getDefaultApiClient();
        String endpoint = FSConfiguration.get(conf, scheme, Constants.ENDPOINT_KEY_SUFFIX, Constants.DEFAULT_CLIENT_ENDPOINT);
        if (endpoint.endsWith(Constants.SEPARATOR)) {
            endpoint = endpoint.substring(0, endpoint.length() - 1);
        }
        apiClient.setBasePath(endpoint);
        apiClient.addDefaultHeader("X-Lakefs-Client", "lakefs-hadoopfs/" + getClass().getPackage().getImplementationVersion());

        String sessionId = FSConfiguration.get(conf, scheme, Constants.SESSION_ID);
        if (sessionId != null) {
            apiClient.addDefaultCookie("sessionId", sessionId);
        }
        return apiClient;
    }

    public ObjectsApi getObjectsApi() {
        return objectsApi;
    }

    public StagingApi getStagingApi() {
        return stagingApi;
    }

    public RepositoriesApi getRepositoriesApi() {
        return repositoriesApi;
    }

    public BranchesApi getBranchesApi() {
        return branchesApi;
    }

    public ConfigApi getConfigApi() {
        return configApi;
    }

    public InternalApi getInternalApi() {
        return internalApi;
    }
}
