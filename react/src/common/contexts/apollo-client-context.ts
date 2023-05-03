import {ApolloClient, ApolloLink, InMemoryCache} from "@apollo/client";
import backend from "../../config/backend";
import {setContext} from "@apollo/client/link/context";
import {selectAuthStatus} from "../../modules/auth/store/auth-selector";
import {store} from "../../store/store";
import moment from "moment/moment";
import {authRefreshToken} from "../../modules/auth/store/auth-slice";
import {Auth} from "../../modules/auth/service/auth";


const authLink = setContext((_, {headers}) => ({
    headers: {...headers}
}));

const isTokenValid = (auth: Auth | null) => {
    if (!auth) {
        return false;
    }

    if (auth.isEmpty()) {
        return false;
    }

    if (!auth.user?.isTokenValid()) {
        return false;
    }

    if (moment(auth.user?.getExpiration()).subtract(1, 'minutes') > moment()) {
        return false;
    }

    return true;
};

const client = new ApolloClient({
    uri: backend.graphEndpoint,
    cache: new InMemoryCache({
        addTypename: false
    }),
    credentials: 'include',
    link: new ApolloLink((operation) => {
            const auth = selectAuthStatus(store.getState());

            if ( auth.refreshToken && !isTokenValid( auth ) ) {
                console.log( 'ciansoda' );
                store.dispatch( authRefreshToken( auth.refreshToken ) );
            }
        })
});

export default client;