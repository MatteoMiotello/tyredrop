import {ApolloClient, InMemoryCache} from "@apollo/client";
import backend from "../../config/backend";


const client = new ApolloClient({
    uri: backend.graphEndpoint,
    cache: new InMemoryCache(),
});

export default client;