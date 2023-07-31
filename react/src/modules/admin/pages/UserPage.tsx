import React from "react";
import {ALL_USERS} from "../../../common/backend/graph/query/users";
import {FetchAllUsersQuery, FetchAllUsersQueryVariables} from "../../../__generated__/graphql";
import UserAdminTable from "../components/User/UserAdminTable";
import Panel from "../../../common/components-library/Panel";
import {useQuery} from "../../../common/backend/graph/hooks";

const UserPage: React.FC = () => {
    const query = useQuery<FetchAllUsersQuery, FetchAllUsersQueryVariables>( ALL_USERS, {
        variables: {
            pagination: {
                limit: 10,
                offset: 0
            }
        }
    } );

    return <main>
        <Panel>
            <Panel.Title>
                Tutti gli utenti
            </Panel.Title>
            {query.data?.users && <UserAdminTable query={query}/>}
        </Panel>
    </main>;
};

export default UserPage;