# Configuration

Manny's `build` subcommand is a wrapper around `Configurator`. Configurators job does all the business logic for Manny.
First base configurations are loaded, followed by stacks, and finally a deployment manifest is output.

### Bases

Bases are the first thing that are loaded, their data structure can be 
[found here](https://github.com/keikoproj/manny/blob/master/configurator/configurator.go#L113-L138). Manny
configs can inherit from eachother using a field called "Base", which is how base configs got their name.

#### Configuration Inheritance Example

Given the following directory structure:

    usw2
    ├── config.yaml
    ├── vpc1
    │   ├── config.yaml
    │   └── privatelink
    │       └── config.yaml
    └── vpc2
        ├── config.yaml
        └── privatelink
            └── config.yaml

Given that:
- `usw2/vpc1/privatelink/config.yaml` has a `base` field of `../config.yaml` and a `test` field of `test`
- `usw2/vpc1/config.yaml` has a `base` field of `../config.yaml` and a `test` field of `test-2`
- `usw2/config.yaml` has a `test` field of `test-3`
- The user performs: `manny build usw2/vpc1/privatelink`

Manny will:
1. Get a list of files in `usw2/vpc1/privatelink/` and find `config.yaml`.
2. Reach `config.yaml` and extract the base. When a base is found Manny will pause unmarshalling data and will read the 
next base.
3. Read `usw2/vpc1/config.yaml` and discover another base, pausing unmarshalling as it did before.
4. Read `usw2/config.yaml`
5. After unmarshal it goes back to the previous config, creates a `MannyConfig`, and stores that struct in a `slice`
6. After all `MannyConfig` objects have been stored in `Configurator.Bases` Manny will apply each `MannyConfig` in 
`Configurator.Bases` to `Configurator.Global` from first to last index

This process gives us a trail to audit failures and is still quite fast given that empty structs do not occupy memory.
That said, this process can be improved.

#### Picking up additional deployments

Manny is capable of discovering higher level deployments. Assuming the same structure as above, and an input of 
`manny build usw2/vpc1` Manny would discover the `privatelink` directory and then create a second deployment that is 
entirely separated from the first.

Manny does this while evaluating files and directories in a provided directory. It then creates a new `Configurator` 
instance and calls `loadBases()` and `loadStacks()`. The deployments are then added to the parent `Configurator.Stacks`.
