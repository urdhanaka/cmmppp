{
  description = "CMP: Go backend, React/Vite frontend";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux" "aarch64-linux" "aarch64-darwin" "x86_64-darwin"];
      perSystem = {
        config,
        self',
        inputs',
        pkgs,
        system,
        ...
      }: {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            # backend
            go_1_27
            gopls
            goose
            golangci-lint
            go-swag
            delve

            # frontend
            nodejs_24
            pnpm

            # database
            postgresql_18

            # cloud CLIs, useful for local testing against AWS/GCP/Azure
            # AWS and GCP is not yet started
            # awscli2
            # google-cloud-sdk
            azure-cli
          ];
        };
      };
      flake = {};
    };
}
