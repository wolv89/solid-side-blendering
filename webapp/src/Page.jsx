import { createSignal, createEffect, createResource, Show } from 'solid-js';
import { useParams, useLocation } from '@solidjs/router';

const fetchEndpoint = async (pagename) => {
    const response = await fetch(`http://blewb.build:8338/content/${pagename}/`)
    return response.text()
}

function Page() {

    const params = useParams()

    const [pagename, setPagename] = createSignal()
    const [content] = createResource(pagename, fetchEndpoint)

    createEffect(() => {
        if (!params.page) {
            setPagename("home")
        } else {
            setPagename(params.page)
        }
    })

    return (
        <div>
            <Show when={content.loading}>
                <p>Loading...</p>
            </Show>
            <Show when={content()}>
                <div innerHTML={content()} />
            </Show>
        </div>
    );
}

export default Page;
