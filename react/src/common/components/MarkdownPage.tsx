import React, {memo, useEffect, useState} from "react";
import Spinner from "./Spinner";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";

type MarkdownPageProps = {
    path: string
}

const MarkdownPage  = memo( ({path}: MarkdownPageProps ) => {
    const [ content, setContent ] = useState<string | undefined>();
    const [loading, setLoading] = useState(false);

    useEffect( () => {
        import(`./../../assets/pages/${path}.md`).then(res => {
            setLoading(true);

            fetch(res.default)
                .then(response => response.text())
                .then(text => setContent(text))
                .finally( () => setLoading(false) );
        });
    }, [] );

    return <div className="relative">
        {loading && <Spinner/>}
        {content && <ReactMarkdown className="markdown" remarkPlugins={[remarkGfm]}>{content}</ReactMarkdown>}
    </div>;
});

MarkdownPage.displayName = 'MarkdownPage';

export default MarkdownPage;