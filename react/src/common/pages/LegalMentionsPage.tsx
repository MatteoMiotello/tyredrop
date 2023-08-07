import React from "react";
import Panel from "../components-library/Panel";
import MarkdownPage from "../components/MarkdownPage";

const LegalMentionsPage: React.FC = () => {
    return <main>
        <Panel>
            <MarkdownPage path="legal-mentions"/>
        </Panel>
    </main>;
};

export default LegalMentionsPage;