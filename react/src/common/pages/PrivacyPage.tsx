import React from "react";
import Panel from "../components-library/Panel";
import MarkdownPage from "../components/MarkdownPage";

const PrivacyPage: React.FC = () => {
    return <main>
        <Panel className="px-8">
            <MarkdownPage path="privacy"/>
        </Panel>
    </main>;
};
export default PrivacyPage;