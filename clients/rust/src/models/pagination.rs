/*
 * lakeFS API
 *
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: services@treeverse.io
 * Generated by: https://openapi-generator.tech
 */

use crate::models;

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct Pagination {
    /// Next page is available
    #[serde(rename = "has_more")]
    pub has_more: bool,
    /// Token used to retrieve the next page
    #[serde(rename = "next_offset")]
    pub next_offset: String,
    /// Number of values found in the results
    #[serde(rename = "results")]
    pub results: i32,
    /// Maximal number of entries per page
    #[serde(rename = "max_per_page")]
    pub max_per_page: i32,
}

impl Pagination {
    pub fn new(has_more: bool, next_offset: String, results: i32, max_per_page: i32) -> Pagination {
        Pagination {
            has_more,
            next_offset,
            results,
            max_per_page,
        }
    }
}

