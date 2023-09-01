import React from "react";
import Panel from "../components-library/Panel";
import Markdown from "../components/MarkdownPage";


const GeneralTermsPage: React.FC = () => {
    return <main className="flex items-center">
        <Panel className="px-8">
            <Markdown path="terms"/>
        </Panel>
    </main>;
};

export default GeneralTermsPage;