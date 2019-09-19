# TCG-SITEMAP

This program generates a series of gzipped sitemaps and a sitemap index which references those sitemaps, from a single CSV file. The CSV file is expected to be located in the root directory and, in this example, is formatted in the following way:

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
As currently configured, this tool will ignore **,{AVAILABLE_RECORDS}** since this bit of information is irrelevant in this example (the URLs we intend to generate only rely on **{LASTNAME}** AND **{FIRSTNAME}**).

Each inidvidual sitemap will contain a max of 49,999 entires.

## Configuration

Create a file in the root directory called **config.yaml**. The file should be structured as follows:

```
baseURL       : https://{yourDomain}
sitemapDir    : sitemaps
sitemapPrefix : lol-
sitemapSuffix : -sitemap.xml
```

**Note:** In this example, the value of `sitemapSuffix` is "-sitemap.xml". However, the gzipping process will produce files and references to those files with suffixes such as "-sitemap.xml.gz".

A configuration file such as this will produce a sitemap index named **sitemap.xml** structured like this:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <sitemap>
    <loc>lol-0-sitemap.xml</loc>
  </sitemap>
  <sitemap>
    <loc>lol-1-sitemap.xml</loc>
  </sitemap>
  <sitemap>
    <loc>lol-2-sitemap.xml</loc>
  </sitemap>
  ...
</sitemapindex>
```

... and individual sitemaps (each named the same as the references above). It will be structured like this:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>https://{yourDomain}/{firstname}-{lastname}/</loc>
  </url>
  <url>
    <loc>https://{yourDomain}/{firstname}-{lastname}/</loc>
  </url>
  <url>
    <loc>https://{yourDomain}/{firstname}-{lastname}/</loc>
  </url>
  ...
</urlset>
```

## Usage

Compile with `go build` and then run `./tcg-sitemap` in the project directory. Or, simply run `go run main.go {name-of-CSV-file}` in the project directory. The sitemap index and individual sitemaps will appear in the folder specified in **config.yaml**.