import * as React from 'react';
import { Item, Icon } from 'semantic-ui-react';
import { Story } from './proto/hackernews_pb';

type StoryListProps = {
  stories: Story.AsObject[],
  selected: Story.AsObject | null,
  onStorySelect: (id: number) => void
};

const StoryList: React.SFC<StoryListProps> = (props) => {
  return (
    <Item.Group divided={true}>
      {props.stories.map((story, i) =>
        <Item
          style={props.selected && story.id === props.selected.id
            ? {'backgroundColor': 'rgba(0, 0, 0, 0.08)'}
            : {}
          }
          key={i}
          onClick={() => {
            if (story.id) {
              props.onStorySelect(story.id);
            }
          }}
        >
          <Item.Content

          >
            <Item.Header as="a">{story.title}</Item.Header>
            <Item.Extra><Icon  name="star" />{story.score} | <Icon  name="user" />{story.by}</Item.Extra>
          </Item.Content>
        </Item>
      )}
    </Item.Group>
  );
};

export default StoryList;
