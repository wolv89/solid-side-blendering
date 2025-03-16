import { A } from "@solidjs/router";

const StandardLayout = (props) => (
    <>
        <header class="header">
            <div class="header-bar">
                <h1 class="logo">
                    <a href="/">LOGO</a>
                </h1>
                <nav class="navigation">
                    <ol>
                        <li><A activeClass="is-active" href="/" end={true}>Home</A></li>
                        <li><A activeClass="is-active" href="/services">Services</A></li>
                        <li><A activeClass="is-active" href="/about">About</A></li>
                        <li><A activeClass="is-active" href="/contact">Contact</A></li>
                    </ol>
                </nav>
            </div>
        </header>
        <section class="main">
            <div class="wrapper">
                {props.children}
            </div>
        </section>
        <footer class="footer">
            <div class="footer-content">
                <h6>Address</h6>
                <p>100 Main Street, Golang 1234, Victoria Australia</p>
                <h6>Opening Hours</h6>
                <p>Mon - Fri: 8am - 4pm<br />Sat: 10am - 2pm<br />Sun: Closed</p>
                <p class="attribution">Site and contents &copy; 2025</p>
            </div>
        </footer>
    </>
);


export default StandardLayout;
