export default () => (
<ingress>
        <metadata name="my-ingress">
            <annotation key="nginx.ingress.kubernetes.io/app-root">/home</annotation>
            <annotation key="nginx.ingress.kubernetes.io/auth-realm">/home</annotation>
        </metadata>
    </ingress>
)