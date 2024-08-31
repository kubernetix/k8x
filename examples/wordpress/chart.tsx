const replicas = Number(process.env["VARIABLE"]);

export default () => (
  <cluster config="~/.kube/config">
    <namespace name="default">
      <deployment>
        <spec replicas={replicas}>
          <selector>
            <match-label key="appp">snowflake</match-label>
            <match-expression key="component" operator="In" values={["cache"]}>
              redis
            </match-expression>
          </selector>
        </spec>

        <template>
          <spec>
            <container
              image="registry.k8s.io/serve_hostname"
              imagePullPolicy="Always"
              name="snowflake"
            ></container>
          </spec>
        </template>
      </deployment>

      <service>
        <spec>
          <selector>
            <match-label key="app.kubernetes.io/name">MyApp</match-label>
          </selector>

          <port protocol="TCP" port={80} targetPort={80} />
        </spec>
      </service>
    </namespace>
  </cluster>
);
