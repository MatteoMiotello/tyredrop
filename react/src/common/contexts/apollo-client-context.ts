import {ApolloClient, InMemoryCache} from "@apollo/client";
import backend from "../../config/backend";


const client = new ApolloClient({
    uri: backend.graphEndpoint,
    cache: new InMemoryCache({
        addTypename: false
    }),
    credentials: 'include'
});

export default client;