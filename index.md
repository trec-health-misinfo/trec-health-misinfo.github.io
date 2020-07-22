---
layout: default
---

# TREC Health Misinformation Track (2020)

## Track Overview
Misinformation represents a key problem when using search engines to guide any decision-making task: Are users able to discern authoritative from unreliable information and correct from incorrect information? This problem is further exacerbated when the search occurs within uncontrolled data collections, such as the web, where information can be unreliable, misleading, highly technical, and can lead to unfounded escalations. Information from search-engine results can significantly influence decisions, and research shows that increasing the amount of incorrect information about a topic presented in a Search Engine Result Page (SERP) can impel users to make incorrect decisions.

In this context, the TREC 2020 Misinformation track fosters research on retrieval methods that promote reliable and correct information over misinformation. The track offers the following tasks:
* **Total Recall Task**: The goal is to identify all the documents conveying incorrect information for a specific set of topics;
* **Ad-hoc Retrieval Task**: The goal is to design a ranking model that promotes credible and correct information over incorrect information;
* **Evaluation Meta Task**: The goal is to develop new evaluation methods that reflect the credibility and correctness of documents, as well as traditional relevance.

This year, we have focused the track specifically on misinformation related to COVID-19 and SARS-CoV-2, adopting a news corpus from January to April, 2020 as the basis for our test collection. As our understanding of the disease evolved over this period some facts became better known. For example, at one point, it was suggested that Ibuprofen might worsen COVID-19. A retrieval effort undertaken today should avoid returning these articles, or else label them as potentially misleading.

<!--#### These guidelines are still in draft form. We invite comments and suggested changes from participants. We plan to finalize the guidelines on June 5, 2020.-->

## Tentative Schedule
* **~~June 30, 2020~~ ~~July 17, 2020~~** [Topics are released](topics.xml);
* **September 1, 2020** Runs due:
  * Before 7am September 1, 2020 Eastern Time Zone (Gaithersburg, MD, USA);
  * Please check TREC active participants page for the runs submission link;
* **October 2020** Results returned;
* **October 2020** Notebook paper due;
* **November 18-20, 2020** TREC Conference;
* **February 2021** Final report due.

## Collection

#### Corpus
For the TREC Health Misinformation 2020 track, we will be using the documents found in [CommonCrawl News crawl](https://commoncrawl.org/2016/10/news-dataset-available/) from January, 1st 2020 to April 30th, 2020. CommonCrawl News contains news articles from news sites all over the world.  
The format of the collection follows a standard Web ARChive (WARC) format. Each document in a WARC file contains a WARC header and the raw data from the crawl. To learn more about the format of the collection and examples of the full WARC extract, please see the CommonCrawl website [here](https://commoncrawl.org/the-data/get-started/).

The corpus contains non-English documents. Non-English documents are not relevant, even if the document would be relevant the non-English language.

**Instructions on how to download the collection:** The CC News Crawl is available on AWS S3. You will need the [AWS CLI](https://aws.amazon.com/cli/) to download it.
In the following, you can find the commands to download the data for all four months (please replace the destination with the intended destination in your machine).

```
$ aws --no-sign-request s3 sync s3://commoncrawl/crawl-data/CC-NEWS/2020/01 /path/to/local/destination
$ aws --no-sign-request s3 sync s3://commoncrawl/crawl-data/CC-NEWS/2020/02 /path/to/local/destination
$ aws --no-sign-request s3 sync s3://commoncrawl/crawl-data/CC-NEWS/2020/03 /path/to/local/destination
$ aws --no-sign-request s3 sync s3://commoncrawl/crawl-data/CC-NEWS/2020/04 /path/to/local/destination
```

**Optional: How to get WET format.** Common Crawl also informally provides a tool to get the text extracts (WET format). WET files contain the extracted plain text with tags (HTML, scripts, etc) removed. Unless you have a reason to do otherwise, we recommend working with these text extracts. If you would like to obtain the WET format for the news crawl, please see the instructions [here](https://groups.google.com/d/msg/common-crawl/hsb90GHq6to/SSVocyq8AAAJ). More information on the WARC and WET file formats can be found [here](https://commoncrawl.org/the-data/get-started/).

**Document Identifier:** The WARC header for each document in a WARC file contains a "WARC-Record-ID" field. For our purposes, the value of the WARC-Record-ID field is considered the document identifier (the "docno").  If you plan to use WET files, please use the WARC-Refers-To field instead.

#### Topics
The track focuses on topics within the consumer health search domain (people seeking health advice online). For TREC 2020 the track will focus on COVID-19. The recent coronavirus crisis represents a good example of uncontrolled proliferation of misinformation, which can have serious consequences on consumer health.

Unlike previous tracks, the assessors will not be creating their own topic statements. Instead, the assessors will be provided with topics that include title, description, answer, narrative, and evidence fields. The title field of each topic is built as a pair of treatment and disease, where for TREC 2020, the disease is always COVID-19. The description is in the form of a question and is built as a triplet of (treatment, effect, disease) where the effect can be: cause, prevent, worsen, cure, help. Only these terms will be used, so that descriptions are all of the form: "Can X Y COVID-19?", where X is a treatment and Y is one of the five effect terms.

The answer field  is one of "yes" or "no". *You should assume that this field specifies the correct answer for the purposes of this task.* This answer corresponds to the topic writer's best understanding of medical consensus at the time of topic creation, but it is not medical advice, and should not be taken as truth outside of the context of this track. The evidence field contains the URL of a page from the open Web that was used to determine this answer. This page may or may not be part of the corpus.

For the total recall task, participants should identify documents contradicting the answer. For the adhoc task, participants should return the most credible and complete information supporting the answer. Note that for many topics the corpus may contain a large number of documents that would be relevant in the traditional topical sense, but which neither support nor contradict the answer.

The topics will be provided as `XML` files using the following format:

```xml
<topics>
<topic>
<number>0</number>
<title>ibuprofen COVID-19</title>
<description>Can ibuprofen worsen COVID-19?</description>
<answer>no</answer>
<evidence>https://www.ncbi.nlm.nih.gov/pmc/articles/PMC7287029</evidence>
<narrative>Ibuprofen is an anti-inflammatory drug used to reduce fever and treat pain and
inflammation. Recently, there has been a debate over whether Ibuprofen can worsen 
the effects of COVID-19. A helpful document might explain in clear language that
there is no scientific evidence supporting this concern. A harmful document might
create anxiety and/or cause people to avoid taking the drug.</narrative>
</topic>
<topic>
...
</topic>
</topics>
```

The narrative and evidence fields are intended to aid with assessment and should not be used for automatic runs. All of the other fields may be used by automatic runs.

## Runs
For the total recall and adhoc tasks, runs may be either automatic or manual. An automatic run is made without any tuning or manual influence. Best practice for an automatic run is to avoid using the topics or even looking at them until all decisions and code have been written to produce the automatic run. The narrative field and evidence field of topics should not be used for automatic runs, but all other topic fields may be used.  

A manual run is anything that is not an automatic run. Manual runs commonly have some human input based on the topics, e.g., hand-crafted queries or relevance feedback. The narrative and evidence fields may be used for manual runs, but use of these fields makes the run a manual run, even if all other processing is automatic.

Submission format will follow the standard TREC run format as follows:

```
qid Q0 docno rank score tag
```
where:
* `qid`: the topic number;
* `Q0`: unused and should always be Q0;
* `docno`: the official document id number returned by your system for the topic `qid`;
* `ran`: the rank the document is retrieved;
* `score`: the score (integer or floating point) that generated the ranking. The score must be in descending (non-increasing) order. The score is important to handle tied scores. (`trec_eval` sorts documents by the specified scores values and not your ranks values);
* `tag`: a tag that uniquely identifies your group AND the method you used to produce the run. Each run should have a different tag.

The fields should be spectated with a whitespace. The width of the columns in the format is not important, but it is important to include all columns and have some amount of white space between the columns.
Example run is shown below:
```
1 Q0 doc-1018wb-57-17875 1 14.8928003311 myGroupNameMyMethodName
1 Q0 doc-1311wb-18-06089 2 14.7590999603 myGroupNameMyMethodName
1 Q0 doc-1113wb-13-13513 3 14.5707998276 myGroupNameMyMethodName
1 Q0 doc-1200wb-47-01250 4 14.5642995834 myGroupNameMyMethodName
1 Q0 doc-0205wb-37-09270 5 14.3723001481 myGroupNameMyMethodName
...
```

## Tasks

#### Task 1 - Total Recall

**Task Description:** Documents contradicting the topic's answer are assumed to be misinformation. Participants must identify all documents in a collection that promulgate, promote, and/or support that misinformation. For example, for the example topic above ("Can Ibuprofen worsen COVID-19?"), you must identify all documents indicating that Ibuprofen can worsen COVID-19. Documents making this claim for the purposes of debunking it are not misinformation.

**Runs:** Runs should rank documents according to the likelihood that they promulgate misinformation. Submission format will follow the standard TREC run format, as specified above. You may submit up to three runs of up to 10,000 ranked documents for each topic.

**Evaluation:** Runs will be compared using *gain curves*, which plots recall as a function of rank. The primary metric is R-precision, or equivalently, R-recall, the recall achieved at rank R, where R is the number of positively labeled documents in the collection.

#### Task 2 - AdHoc Retrieval

**Task Description:** Participants devise search technologies that promote credible and correct information over incorrect information, with the assumption that correct information can better lead people to make correct decisions.

Given the corpus and topics, your task is to return relevant, credible, and correct information that will help searchers make correct decisions. 
You should assume that the statement included in the topic description is correct or not, based on the answer field, even if you know current medical or other evidence suggests otherwise.

Note that this task is more than simply a new definition of what is relevant. There are three types of results: correct and relevant, incorrect, and non-relevant. It is important that search results avoid containing incorrect results, and ranking non-relevant results above incorrect is preferred. In place of notions of correctness, the credibility of the information source is useful, and relevant and credible information is preferred.

**Runs:** Submission format will follow the standard TREC run format, as specified above. For each topic, please return 1,000 ranked documents.  
Participating groups will be allowed to submit as many runs as they like, but they need authorization from the Track organizers before submitting more than 10 runs. Not all runs are likely to be used for pooling and groups will likely need to specify a preference ordering for pooling purposes.

**Evaluation:** The final qrels will contain assessments with respect to the following criteria:
* *Relevance*: whether the document is relevant to the topic;
* *Credibility*: whether the document is considered credible by the assessor;
* *Correctness*: whether the document contains correct information with respect to the answer provided in the topic description.
Note that non-relevant documents will not be assessed with respect to credibility and correctness. 

Submitted runs will be evaluated with respect to the three criteria: relevance, credibility, and correctness. We will design specific measures to account for those aspects and to penalize systems which retrieve incorrect documents.

We will also evaluate runs in terms of traditional relevance measures, e.g., nDCG@10 and MAP, with the goal of comparing performance measures between the relevance-only measures and the measures that combine relevance, credibility, and correctness.

#### Task 3 - Evaluation Meta Task
Details to be announced later.

## Organizers

* [Charles Clarke, University of Waterloo](https://cs.uwaterloo.ca/about/people/claclark)
* [Maria Maistro, University of Copenhagen](https://di.ku.dk/english/staff/?pure=en/persons/641366)
* [Mark Smucker, University of Waterloo](http://mansci.uwaterloo.ca/~msmucker/)
* [Guido Zuccon, University of Queensland](http://www.zuccon.net/)


## Contact
For more information or to ask questions, join the [google groups](https://groups.google.com/forum/#!forum/trec-health-misinformation-track)
