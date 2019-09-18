# TCG-SITEMAP

This program generates a series of sitemaps and a sitemap index which references those sitemaps, from a single CSV file. The CSV file is expected to be formatted in the following way:

```
{LASTNAME},{FIRSTNAME},{AVAILABLE_RECORDS}
{LASTNAME},{FIRSTNAME},{AVAILABLE_RECORDS}
{LASTNAME},{FIRSTNAME},{AVAILABLE_RECORDS}
...
```

For example:

```
SMITH,LUCY,90
MARTINEZ,ANDREA,70
PATEL,IRENE,24
...
```

Each inidvidual sitemap will contain a max of 49,000 entires.

## Usage

After compiling, run `./tcg-sitemap` in the project directory. The sitemap index and sitemaps will appear in a folder named **/sitemaps**.