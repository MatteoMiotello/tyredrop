import {faCopyright, faEuro, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import Panel from "../../../common/components-library/Panel";
import {Stat} from "../../../common/components/shelly-ui";
import {useQuery} from "../../../common/backend/graph/hooks";
import {STATS} from "../../../common/backend/graph/query/stat";
import Spinner from "../../../common/components/Spinner";
import {AllOrdersQuery, AllOrdersQueryVariables} from "../../../__generated__/graphql";
import {FETCH_ALL_ORDERS} from "../../../common/backend/graph/query/order";
import OrderAdminTable from "../components/Order/OrderAdminTable";

const HomePage: React.FC = () => {
    const statsQuery = useQuery(STATS);
    const query = useQuery<AllOrdersQuery, AllOrdersQueryVariables>(FETCH_ALL_ORDERS, {
        variables: {
            pagination: {
                offset: 0,
                limit: 10
            },
            ordering: [
                {
                    column: 'created_at',
                    desc: true
                },
            ]
        }
    });

    return <main>
        <div className=" p-4 lg:p-24 flex flex-col lg:flex-row justify-around gap-4">
            {
                statsQuery.loading && <Spinner/>
            }
            {
                statsQuery.data &&
                <>
                    <Stat>
                        <Stat.Figure className="text-secondary">
                            <FontAwesomeIcon icon={faCopyright}/>
                        </Stat.Figure>
                        <Stat.Title>
                            Brand
                        </Stat.Title>
                        <Stat.Value>
                            {statsQuery.data.stats.bestBrand.name}
                        </Stat.Value>
                        <Stat.Desc>
                            Il brand con più prodotti venduti
                        </Stat.Desc>
                    </Stat>
                    <Stat>
                        <Stat.Figure className="text-primary">
                            <FontAwesomeIcon icon={faEuro}/>
                        </Stat.Figure>
                        <Stat.Title>
                            Ordini
                        </Stat.Title>
                        <Stat.Value>
                            {statsQuery.data.stats.totalOrders}
                        </Stat.Value>
                        <Stat.Desc>
                            Totale ordini evasi nell'ultimo mese
                        </Stat.Desc>
                    </Stat>
                    <Stat>
                        <Stat.Figure className="text-accent">
                            <FontAwesomeIcon icon={faUser}/>
                        </Stat.Figure>
                        <Stat.Title>
                            Clienti
                        </Stat.Title>
                        <Stat.Value>
                            {statsQuery.data.stats.totalUsers}
                        </Stat.Value>
                        <Stat.Desc>
                            Totale clienti attivi attuali
                        </Stat.Desc>
                    </Stat>
                </>
            }
        </div>
        <Panel>
            <Panel.Title>
                Ultimi ordini
            </Panel.Title>
            {
                query.data &&
                <OrderAdminTable query={query}/>
            }
        </Panel>
    </main>
        ;
};

export default HomePage;