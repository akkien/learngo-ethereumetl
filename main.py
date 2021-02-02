def get_partitions(start_block, end_block, partition_batch_size):
    """Yield partitions based on input data type."""
    for batch_start_block in range(start_block, end_block + 1, partition_batch_size):
        batch_end_block = batch_start_block + partition_batch_size - 1
        if batch_end_block > end_block:
            batch_end_block = end_block
        yield batch_start_block, batch_end_block


partitions = list(get_partitions(5, 20, 5))
print(partitions)
