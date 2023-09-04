import React from "react";
import {useLoaderData} from "react-router-dom";
import {
    AllUserInvoicesQuery,
    AllUserInvoicesQueryVariables,
    FetchUserQuery,
    UserBilling
} from "../../__generated__/graphql";
import {useQuery} from "../../common/backend/graph/hooks";
import { ALL_USER_INVOICES} from "../../common/backend/graph/query/invoice";
import Panel from "../../common/components-library/Panel";
import InvoiceTable from "../invoice/components/InvoiceTable";

const UserInvoicesPage: React.FC = () => {
    const user = useLoaderData() as { data: FetchUserQuery, loading: boolean };

    if ( !user.data.user?.userBilling ) {
        return <main className="text-2xl p-24"> Non e` presente alcuna fattura </main>;
    }

    const query = useQuery<AllUserInvoicesQuery, AllUserInvoicesQueryVariables>( ALL_USER_INVOICES, {
        variables: {
            pagination: {
                limit: 20,
                offset: 0
            },
            input: {
                userBillingId: user.data.user.userBilling.id
            }
        }
    } );


    return <main>
        <Panel>
            <Panel.Title>
                Tutte le fatture
            </Panel.Title>
            {
                query.data &&
                <InvoiceTable query={query} userBilling={user.data.user.userBilling as UserBilling}/>
            }
        </Panel>
    </main>;
};

export default UserInvoicesPage;