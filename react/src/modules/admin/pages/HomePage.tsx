import {faEuro, faRuler, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import Panel from "../../../common/components-library/Panel";
import {Stat} from "../../../common/components/shelly-ui";

const HomePage: React.FC = () => {
    return <main>
        <div className="p-24 flex justify-around gap-24">
            <Stat>
                <Stat.Figure className="text-secondary">
                    <FontAwesomeIcon icon={faRuler}/>
                </Stat.Figure>
                <Stat.Title>
                  Miglior misura
                </Stat.Title>
                <Stat.Value>
                    255/15 R7
                </Stat.Value>
                <Stat.Desc>
                    Misura di pneumatici piu` venduta
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
                    15.000€
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
                    1000
                </Stat.Value>
                <Stat.Desc>
                    Totale clienti attivi attuali
                </Stat.Desc>
            </Stat>
        </div>
        <Panel>
            <Panel.Title>
                Ultimi ordini
            </Panel.Title>
            
        </Panel>
    </main>;
};

export default HomePage;