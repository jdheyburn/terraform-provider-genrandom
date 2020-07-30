# terraform-provider-genrandom

TF Provider to generate a random number.

## Local Development

To build:

```bash
make build
```

To test:

```bash
make test
```

To copy locally for testing:

```bash
cp ~/go/bin/terraform-provider-genrandom ~/.terraform.d/plugins/
```

## Resources

### genrandom_int

Example usage:

```hcl
resource "genrandom_int" "rando" {
    min = 5
    max = 10
}

output "my_random_int" {
    value = "The random int is: ${genrandom_int.rando.value}"
}
```
