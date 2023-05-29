import {
    ApolloClient,
    ApolloLink,
    FetchResult,
    HttpLink,
    InMemoryCache,
    NextLink,
    Observable,
    Operation, from, fromPromise
} from "@apollo/client";
import { onError} from "@apollo/client/link/error";
import {GraphQLError} from "graphql/error";
import backend from "../../config/backend";
import moment from "moment/moment";
import {selectAuthStatus} from "../../modules/auth/store/auth-selector";
import {authRefreshToken} from "../../modules/auth/store/auth-slice";
import {store} from "../../store/store";

const httpLink = new HttpLink({
    uri: backend.graphEndpoint,
    credentials: 'include'
});


const refreshTokenLink = new ApolloLink((operation: Operation, forward: NextLink): Observable<FetchResult> | Observable<{
    data?: Record<string, any> | null;
    context?: Record<string, any>;
    errors?: ReadonlyArray<GraphQLError>;
    extensions?: Record<string, any>
}> => {
    const auth = selectAuthStatus(store.getState());

    if ( auth.isEmpty() || auth.isPending() ) {
        return forward(operation);
    }

    if (auth && !auth.user?.isTokenValid()) {
        return fromPromise(store.dispatch(authRefreshToken(auth.refreshToken)))
            .flatMap(res => {
                return forward(operation);
            });
    }

    const user = auth?.user;

    if (user && user.getExpiration() !== null) {
        if (user.getExpiration() as Date >= moment().subtract(1, 'minutes').toDate()) {
            return fromPromise(store.dispatch(authRefreshToken(auth.refreshToken)))
                .flatMap(res => {
                    return forward(operation);
                });
        }
    }

    return forward(operation);
});


const errorLink = onError(
    ({graphQLErrors, networkError, operation, forward}) => {
        const auth = selectAuthStatus(store.getState());

        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        if (networkError && networkError?.statusCode == 401 && auth?.refreshToken) {
            return fromPromise(store.dispatch(authRefreshToken(auth.refreshToken)))
                .flatMap(res => {
                    return forward(operation);
                });
        }
    }
);

const client = new ApolloClient({
    cache: new InMemoryCache({
        addTypename: true
    }),
    link: from([errorLink, refreshTokenLink, httpLink])
});

export default client;