---
layout: default
---

# TREC Health Misinformation Track (2021)

**Update**: The 2021 track is still being designed. Keep an eye on our website for the latest updates and make sure you join the [google groups](https://groups.google.com/forum/#!forum/trec-health-misinformation-track)! 

## Track Introduction
Misinformation represents a key problem when using search engines to guide any decision-making task: Are users able to discern authoritative from unreliable information and correct from incorrect information? This problem is further exacerbated when the search occurs within uncontrolled data collections, such as the web, where information can be unreliable, misleading, highly technical, and can lead to unfounded escalations. Information from search-engine results can significantly influence decisions, and research shows that increasing the amount of incorrect information about a topic presented in a Search Engine Result Page (SERP) can impel users to make incorrect decisions.

In this context, the TREC Health Misinformation track fosters research on retrieval methods that promote reliable and correct information over misinformation.

The 2021 track is still being designed. Please bookmark this page to keep up to date as we develop the track for this year's TREC.

<!-- This year, we have focused the track ... -->

<!--#### These guidelines are still in draft form. We invite comments and suggested changes from participants. We plan to finalize the guidelines on June 5, 2020.-->

## Tentative Schedule
* **June-July 2021** Topics will be released;
* **September, 2021** Runs due;
* **September 30 2021** Results returned;
* **October 2021** Notebook paper due;
* **November 17-19, 2021** TREC Conference;
* **February 2022** Final report due.


## Collection

#### Corpus
This year we wil be using the uncleaned version of the uncleaned C4 dataset used by Google to train their T5 model. The collection is comprised of text extracts from the April 2019 snapshot of Common Crawl. The Collection contains ~1B english documents.

You can download the corpus on a Debian/Ubuntu machine using the following commands.
```
sudo apt-get install git-lfs 
git lfs install
GIT_LFS_SKIP_SMUDGE=1 git clone https://huggingface.co/datasets/allenai/c4
cd c4
git lfs pull --include="en.noclean/c4-train*"
```
The collection is made up of the 7168 gzipped jsonl files located in the en.nolean directory. Each file containes ~150k documents, one in each line. A document is a json object with the fields `text`, `url` and `timestamp`. They do not contain a document identifier. The document identifier for each document will be the "c4nc-<file#>-<line#>. <file#> is 4 digits and <line#> is 5 digits. <line#> start at zero for each file and <file#> can be derived from the file names. For example, the document on the second line of file c4-train.01234-of-07168.json.gz would have a docno of c4nc-1234-00001.

One way to insert document identifiers is by using the [provided script](renamer.go). Another would be name the documents as you index them.

#### Topics

Topics are still being created.

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

**Task Description:** Documents contradicting the topic's answer are assumed to be misinformation. Participants must identify all documents in a collection that promulgate, promote, and/or support that misinformation. For example, for the example topic "Can aspirin improve the lives of people with vascular dementia?", you must identify all documents indicating that Asprin can help to improve the vascular system and benefit people with dementia. Documents making this claim for the purposes of debunking it are not misinformation.

**Runs:** Runs should rank documents according to the likelihood that they promulgate misinformation. Submission format will follow the standard TREC run format, as specified above. You may submit up to three runs of up to 10,000 ranked documents for each topic.

**Evaluation:** Runs will be compared using *gain curves*, which plots recall as a function of rank. The primary metric is R-precision, or equivalently, R-recall, the recall achieved at rank R, where R is the number of positively labeled documents in the collection.

#### Task 2 - AdHoc Retrieval

**Task Description:** Participants devise search technologies that promote credible and correct information over incorrect information, with the assumption that correct information can better lead people to make correct decisions.

Given the corpus and topics, your task is to return relevant, credible, and correct information that will help searchers make correct decisions. 
You should assume that the statement included in the topic description is correct or not, based on the answer field, even if you know current medical or other evidence suggests otherwise.

Note that this task is more than simply a new definition of what is relevant. There are three types of results: correct and relevant, incorrect, and non-relevant. It is important that search results avoid containing incorrect results, and ranking non-relevant results above incorrect is preferred. In place of notions of correctness, the credibility of the information source is useful, and relevant and credible information is preferred.

**Runs:** Submission format will follow the standard TREC run format, as specified above. For each topic, please return 1,000 ranked documents.  
Participating groups will be allowed to submit as many runs as they like, but they need authorization from the Track organizers before submitting more than 10 runs. Not all runs are likely to be used for pooling and groups will likely need to specify a preference ordering for pooling purposes.

**Evaluation:**
The final qrels will contain assessments with respect to the following criteria:
* *Usefulness*: whether the document document contains information that a search user would find useful in answering the topic's question;  
* *Correctness*: whether the document contains a definitive and correct answer to the topic's question;
* *Credibility*: whether the document is considered credible by the assessor;
Note that not-useful documents will not be assessed with respect to credibility and correctness. 

Submitted runs will be evaluated with respect to the three criteria: usefulness, correctness, and credibility. We will design specific measures to account for those aspects and to penalize systems which retrieve incorrect documents.

We will also evaluate runs in terms of traditional relevance measures, e.g., nDCG@10 and MAP, with the goal of comparing performance measures between the relevance-only measures and the measures that combine usefulness, credibility, and correctness.  

#### Task 3 - Evaluation Meta Task
Details forthcoming.


## Organizers

* [Charles Clarke, University of Waterloo](https://cs.uwaterloo.ca/about/people/claclark)
* [Maria Maistro, University of Copenhagen](https://di.ku.dk/english/staff/?pure=en/persons/641366)
* [Mark Smucker, University of Waterloo](http://mansci.uwaterloo.ca/~msmucker/)


## Contact
For more information or to ask questions, join the [google groups](https://groups.google.com/forum/#!forum/trec-health-misinformation-track)
